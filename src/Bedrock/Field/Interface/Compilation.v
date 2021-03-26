Require Import Rupicola.Lib.Api.
Require Import Crypto.Bedrock.Specs.Field.
Require Import Crypto.Arithmetic.PrimeFieldTheorems.
Local Open Scope Z_scope.

Section Compile.
  Context {semantics : Semantics.parameters}
          {semantics_ok : Semantics.parameters_ok semantics}.
  Context {field_parameters : FieldParameters}
          {field_representaton : FieldRepresentation}
          {field_representation_ok : FieldRepresentation_ok}.

  (* "Hidden" alias protects a Placeholder (e.g. the pointer reserved for the
     final output value) from having intermediate values stored in it *)
  Definition Hidden := Placeholder.

  Local Ltac prove_field_compilation :=
    repeat straightline';
    handle_call;
    lazymatch goal with
    | |- sep _ _ _ => ecancel_assumption
    | _ => idtac
    end; eauto;
    sepsimpl; repeat straightline'; subst; eauto.

  Lemma compile_binop {name} {op: BinOp name}
        {tr mem locals functions} x y:
    let v := bin_model (feval x) (feval y) in
    forall {P} {pred: P v -> predicate} {k: nlet_eq_k P v} {k_impl}
      Rin Rout x_ptr x_var y_ptr y_var out_ptr out_var,

      (_: spec_of name) functions ->
      bounded_by bin_xbounds x ->
      bounded_by bin_ybounds y ->

      map.get locals out_var = Some out_ptr ->
      (Placeholder out_ptr * Rout)%sep mem ->

      (FElem x_ptr x * FElem y_ptr y * Rin)%sep mem ->
      map.get locals x_var = Some x_ptr ->
      map.get locals y_var = Some y_ptr ->

      (let v := v in
       forall out m (Heq: feval out = v),
         bounded_by bin_outbounds out ->
         sep (FElem out_ptr out) Rout m ->
         (<{ Trace := tr;
             Memory := m;
             Locals := locals;
             Functions := functions }>
          k_impl
          <{ pred (k v eq_refl) }>)) ->
      <{ Trace := tr;
         Memory := mem;
         Locals := locals;
         Functions := functions }>
      cmd.seq
        (cmd.call [] name [expr.var x_var; expr.var y_var;
                          expr.var out_var])
        k_impl
      <{ pred (nlet_eq [out_var] v k) }>.
  Proof. prove_field_compilation. Qed.

  Lemma compile_unop {name} (op: UnOp name) {tr mem locals functions} x:
    let v := un_model (feval x) in
    forall {P} {pred: P v -> predicate} {k: nlet_eq_k P v} {k_impl}
      Rin Rout x_ptr x_var out_ptr out_var,

      (_: spec_of name) functions ->
      bounded_by un_xbounds x ->

      map.get locals out_var = Some out_ptr ->
      (Placeholder out_ptr * Rout)%sep mem ->

      (FElem x_ptr x * Rin)%sep mem ->
      map.get locals x_var = Some x_ptr ->

      (let v := v in
       forall out m,
         bounded_by un_outbounds out ->
         sep (FElem out_ptr out) Rout m ->
         (<{ Trace := tr;
             Memory := m;
             Locals := locals;
             Functions := functions }>
          k_impl
          (* <{ (rew [fun v => P v -> predicate)] Heq in pred) (k (eval out mod M) Heq) }>)) -> *)
          <{ pred (k v eq_refl) }>)) ->
      <{ Trace := tr;
         Memory := mem;
         Locals := locals;
         Functions := functions }>
      cmd.seq
        (cmd.call [] name [expr.var x_var; expr.var out_var])
        k_impl
      <{ pred (nlet_eq [out_var] v k) }>.
  Proof. prove_field_compilation. Qed.

  Ltac cleanup_op_lemma lem := (* This makes [simple apply] work *)
    let lm := fresh in
    let op := match lem with _ _ ?op => op end in
    let op_hd := term_head op in
    let simp proj :=
        (let hd := term_head proj in
         let reduced := (eval cbv [op_hd hd] in proj) in
         change proj with reduced in (type of lm)) in
    pose lem as lm;
    first [ simp (bin_model (BinOp := op));
            simp (bin_xbounds (BinOp := op));
            simp (bin_ybounds (BinOp := op));
            simp (bin_outbounds (BinOp := op))
          | simp (un_model (UnOp := op));
            simp (un_xbounds (UnOp := op));
            simp (un_outbounds (UnOp := op)) ];
    let t := type of lm in
    let t := (eval cbv beta in t) in
    exact (lm: t).

  Notation make_bin_lemma op :=
    ltac:(cleanup_op_lemma (@compile_binop _ op)) (only parsing).

  Definition compile_mul := make_bin_lemma bin_mul.
  Definition compile_add := make_bin_lemma bin_add.
  Definition compile_sub := make_bin_lemma bin_sub.

  Notation make_un_lemma op :=
    ltac:(cleanup_op_lemma (@compile_unop _ op)) (only parsing).

  Definition compile_square := make_un_lemma un_square.
  Definition compile_scmula24 := make_un_lemma un_scmula24.
  Definition compile_inv := make_un_lemma un_inv.

  Lemma compile_felem_copy {tr mem locals functions} x :
    let v := feval x in
    forall {P} {pred: P v -> predicate} {k: nlet_eq_k P v} {k_impl}
      R x_ptr x_var out_ptr out_var,

      spec_of_felem_copy functions ->

      map.get locals out_var = Some out_ptr ->

      (FElem x_ptr x * Placeholder out_ptr * R)%sep mem ->
      map.get locals x_var = Some x_ptr ->

      (let v := v in
       forall m,
         (FElem x_ptr x * FElem out_ptr x * R)%sep m ->
         (<{ Trace := tr;
             Memory := m;
             Locals := locals;
             Functions := functions }>
          k_impl
          <{ pred (k v eq_refl) }>)) ->
      <{ Trace := tr;
         Memory := mem;
         Locals := locals;
         Functions := functions }>
      cmd.seq
        (cmd.call [] felem_copy [expr.var x_var; expr.var out_var])
        k_impl
      <{ pred (nlet_eq [out_var] v k) }>.
  Proof. prove_field_compilation. Qed.

  Lemma compile_felem_small_literal {tr mem locals functions} x:
    let v := x mod M in
    forall {P} {pred: P v -> predicate} {k: nlet_eq_k P v} {k_impl}
      R (wx : word) (out : felem) out_ptr out_var,

      spec_of_felem_small_literal functions ->

      map.get locals out_var = Some out_ptr ->
      (Placeholder out_ptr * R)%sep mem ->

      word.unsigned wx = x ->

      (let v := v in
       forall X m,
         (FElem out_ptr X * R)%sep m ->
         feval X = F.of_Z _ v ->
         bounded_by tight_bounds X ->
         (<{ Trace := tr;
             Memory := m;
             Locals := locals;
             Functions := functions }>
          k_impl
          <{ pred (k v eq_refl) }>)) ->
      <{ Trace := tr;
         Memory := mem;
         Locals := locals;
         Functions := functions }>
      cmd.seq
        (cmd.call [] felem_small_literal
                  [expr.literal x; expr.var out_var])
        k_impl
      <{ pred (nlet_eq [out_var] v k) }>.
  Proof.
    prove_field_compilation.
    match goal with H : _ |- _ =>
                    rewrite word.of_Z_unsigned in H end.
    use_hyp_with_matching_cmd; eauto; [ ].
    subst_lets_in_goal.
    cbv [M]. rewrite <-F.of_Z_mod.
    assumption.
  Qed.
End Compile.

Ltac field_compile_step :=
  first [ simple eapply compile_scmula24 (* must precede compile_mul *)
        | simple eapply compile_mul
        | simple eapply compile_add
        | simple eapply compile_sub
        | simple eapply compile_square
        | simple eapply compile_inv ];
  lazymatch goal with
  | |- feval _ = _ => try eassumption; try reflexivity
  | |- _ => idtac
  end.

(* Change an FElem into a Placeholder to indicate that it is overwritable *)
Ltac free p :=
  match goal with
  | H : sep ?P ?Q ?m |- context [?m] =>
    let x :=
        match type of H with
        | context [FElem p ?x] => x
        end in
    let F :=
        match (eval pattern (FElem p x) in (sep P Q m)) with
        | ?F _ => F end in
    let H' := fresh in
    assert (F (Placeholder p)) as H'
        by (seprewrite FElem_from_bytes; sepsimpl; eexists;
            ecancel_assumption);
    cbv beta in H'; clear H
  end.

(* Protect a pointer (e.g. the pointer reserved for final output) by "hiding"
   the fact that it is available under the "Hidden" alias *)
Ltac hide p :=
  change (Placeholder p) with (Hidden p) in *.
Ltac unhide p :=
  change (Hidden p) with (Placeholder p) in *.
