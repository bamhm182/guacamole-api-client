package types

type GuacSharingProfile struct {
    Name                           string                       `json:"name"`
	Identifier                     string                       `json:"identifier,omitempty"`
	PrimaryConnectionIdentifier    string                       `json:"primaryConnectionIdentifier"`
	Parameters                     GuacSharingProfileParameters `json:"parameters"`
}

type GuacSharingProfileParameters struct {
	ReadOnly    string    `json:"read-only,omitempty"`
}
