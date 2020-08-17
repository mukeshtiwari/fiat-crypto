Require Import bedrock2.Semantics.
Require Import Rupicola.Lib.Api.
Require Import Crypto.Algebra.Group.
Require Import Crypto.Algebra.Hierarchy.
Require Import Crypto.Algebra.ScalarMult.
Require Import Crypto.Arithmetic.PrimeFieldTheorems.
Require Import Crypto.Bedrock.Specs.ScalarField.

Class GroupParameters :=
  { (** mathematical parameters **)
    G : Type;
    eq : G -> G -> Prop;
    add : G -> G -> G;
    zero : G;
    opp : G -> G;
    scalarmult : Z -> G -> G;

    (** function names for bedrock2 **)
    scmul : string;
  }.

Class GroupParameters_ok {group_parameters : GroupParameters} :=
  { group_ok : @group G eq add zero opp;
    scalarmult_ok : @is_scalarmult G eq add zero opp scalarmult;
  }.

Class GroupRepresentation {G : Type} {semantics : Semantics.parameters} :=
  { gelem : Type;
    grepresents : gelem -> G -> Prop;
    GElem : word -> gelem -> Semantics.mem -> Prop;
  }.

Section Specs.
  Context {semantics : Semantics.parameters}
          {scalar_field_parameters : ScalarFieldParameters}
          {scalar_representaton : ScalarRepresentation}.
  Context {group_parameters : GroupParameters}
          {group_representaton : GroupRepresentation (G:=G)}.

  Definition spec_of_scmul : spec_of scmul :=
    (forall! (x old_out : gelem) (k : scalar) (X : G)
           (pout px pk : word),
        (fun Rr mem =>
           grepresents x X
           /\ (exists Ra, (GElem px x * Scalar pk k * Ra)%sep mem)
           /\ (GElem pout old_out * Rr)%sep mem)
          ===>
          scmul @ [pout; px; pk]
          ===>
          (fun _ =>
             liftexists (xk : gelem),
             (emp (grepresents xk (scalarmult (F.to_Z (sceval k)) X))
              * GElem pout xk)%sep)).
End Specs.


