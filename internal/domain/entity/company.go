package entity

type Company struct {
	Inn         string `protobuf:"bytes,1,opt,name=Inn,proto3" json:"Inn,omitempty"`
	Kpp         string `protobuf:"bytes,2,opt,name=Kpp,proto3" json:"Kpp,omitempty"`
	CompanyName string `protobuf:"bytes,3,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	Fio         string `protobuf:"bytes,4,opt,name=Fio,proto3" json:"Fio,omitempty"`
}
