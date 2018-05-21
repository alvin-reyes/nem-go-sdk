package model

type MultisigCosignatoryModification struct {
	ModificationType   int    `json:"modificationType"`
	CosignatoryAccount string `json:"cosignatoryAccount"`
}
