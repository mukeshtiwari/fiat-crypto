(** * The Structure of the [eq] Type *)
(** As described by the project of homotopy type theory, there is a
    lot of structure inherent in the type of propositional equality,
    [eq].  We build up enough lemmas about this structure to deal
    nicely with proofs of equality that come up in practice in this
    project. *)
Require Import Crypto.Util.Isomorphism.
Require Import Crypto.Util.HProp.

(** Most of the structure of proofs of equalities fits into what
    mathematicians call a weak ∞-groupoid.  (In fact, one way of
    *defining* a weak ∞-groupoid is via what's called the J-rule,
    which exactly matches the eliminator [eq_rect] for the equality
    type [eq].)  The first level of this groupoid is given by the
    identity [eq_refl], the binary operation [eq_trans], and the
    inverse of [p : x = y] given by [eq_sym p : y = x].  The second
    level, which we provide here, says that [id ∘ p = p = p ∘ id],
    that [(p ∘ q)⁻¹ = q⁻¹ ∘ p⁻¹], that [(p⁻¹)⁻¹ = p], and that [p ∘
    p⁻¹ = p⁻¹ ∘ p = id].  We prove these here, as well as a few lemmas
    about functoriality of functions over equality. *)
Definition concat_1p {A x y} (p : x = y :> A) : eq_trans eq_refl p = p.
Proof. case p; reflexivity. Defined.
Definition concat_p1 {A x y} (p : x = y :> A) : eq_trans p eq_refl = p.
Proof. case p; reflexivity. Defined.
Definition concat_pV {A x y} (p : x = y :> A) : eq_trans p (eq_sym p) = eq_refl.
Proof. case p; reflexivity. Defined.
Definition concat_Vp {A x y} (p : x = y :> A) : eq_trans (eq_sym p) p = eq_refl.
Proof. case p; reflexivity. Defined.
Definition transport_pp {A} {P : A -> Type} {x y z} (p : x = y) (q : y = z) (k : P x)
: eq_rect _ P k _ (eq_trans p q) = eq_rect _ P (eq_rect _ P k _ p) _ q.
Proof. case q; simpl; reflexivity. Defined.
Lemma transport_const {A P x y} (p : x = y :> A) k
: eq_rect _ (fun _ : A => P) k _ p = k.
Proof. case p; reflexivity. Defined.
Lemma ap_const {A B x y} (b : B) (p : x = y :> A)
: f_equal (fun _ => b) p = eq_refl.
Proof. case p; reflexivity. Defined.
Lemma inv_pp {A x y z} (p : x = y :> A) (q : y = z :> A)
: eq_sym (eq_trans p q) = eq_trans (eq_sym q) (eq_sym p).
Proof. case q; case p; reflexivity. Defined.
Lemma inv_V {A x y} (p : x = y :> A)
: eq_sym (eq_sym p) = p.
Proof. case p; reflexivity. Defined.

(** To classify the equalities of a type [A] at a point [a : A], we
    must give a "code" that stands in for the type [a = x] for each
    [x], and a way of "encoding" proofs of equality into the code that
    we claim represents this type.  To prove that the code is good, it
    turns out that it sufficies to prove that it is freely generated
    by the encoding of [eq_refl], i.e., that [{ x : A & a = x }] and
    [{ x : A & code x }] are equivalent, i.e., that [{ x : A & code x
    }] is an hProp (has at most one element).  When this is the case,
    we have fully classified the type of equalities in [A] at [a] up
    to isomorphism.  This lets us replace proofs of equalities in [A]
    with their codes, which are frequently easier to manipulate.  (For
    example, the code for [x = y :> A * B] is [(fst x = fst y) * (snd
    x = snd y)], whence we can destruct the pair.)

    This method of proof, introduced in homotopy type theory, is
    called encode-decode. *)
Section gen.
  Context {A : Type} {x : A} (code : A -> Type)
          (encode : forall y, x = y -> code y)
          {code_hprop : IsHProp { a : A & code a } }.

  Definition decode' {y} (p : code y) : x = y
    := f_equal (@projT1 _ _) (code_hprop (existT _ x (encode x eq_refl)) (existT _ y p)).

  Definition decode {y} (H : code y) : x = y
    := eq_trans (eq_sym (decode' (encode x eq_refl))) (decode' H).

  Definition deencode {y} (H : x = y) : decode (encode y H) = H.
  Proof.
    destruct H.
    unfold decode.
    edestruct (decode' (encode x eq_refl)).
    reflexivity.
  Defined.

  Global Instance isiso_encode {y} : IsIso (encode y).
  Proof.
    exists (@decode y).
    { intro H.
      unfold decode.
      set (yH := (existT _ y H)).
      change H with (projT2 yH).
      change y with (projT1 yH).
      clearbody yH; clear y H.
      destruct (code_hprop (existT _ x (encode x eq_refl)) yH); simpl.
      abstract (rewrite concat_Vp; reflexivity). }
    { intro p.
      apply deencode. }
  Defined.
End gen.

(** We use the encode-decode method to prove that if [forall x y : A,
    x = y], then [forall (x y : A) (p q : x = y), p = q]. *)
(* It's not clear whether this should be in this file, or in HProp.
   But we require [IsHProp] above, so we leave it here for now. *)
Section hprop.
  Context {A} `{IsHProp A}.

  Let hprop_encode {x y : A} (p : x = y) : unit := tt.

  Local Hint Resolve (fun x => @isiso_encode A x (fun _ => unit)) : typeclass_instances.

  Global Instance ishprop_path_hprop : IsHPropRel (@eq A).
  Proof.
    intros x y p q.
    pose proof (is_left_inv (@hprop_encode x y)) as H'.
    rewrite <- (H' p), <- (H' q).
    apply f_equal; apply allpath_hprop.
  Qed.
End hprop.
