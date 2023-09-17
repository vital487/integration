package b1i

const (
	ACTION_INSERT = "Synchronous Insert"
	ACTION_UPDATE = "Synchronous Update"
)

type Payload interface {
	GetKey() string
	GetXML() string
}

type InboundWS struct {
	DB string
	ObjType string `xml:"Type"`
	Action string
	KeyName string
	KeyId string
	SearchKey1 string
	SearchKey2 string
	Payload interface{}
}


// <?xml version="1.0" encoding="utf-8"?>
// <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
//   <soap:Body>
//     <InboundWS>
//   <DB>0010000100</DB>
//   <Type>2</Type>
//       <Action>Synchronous Insert</Action>
//       <KeyName>CardCode</KeyName>
//       <KeyId> </KeyId>
//       <SearchKey1 />
//       <SearchKey2 />
//       <Payload>
//         <BusinessPartners>
//           <row>
//             <CardCode>123</CardCode>
//             <CardName>Test1</CardName>
//             </row>
//           </BusinessPartners>
//         </Payload>
// </InboundWS>
//   </soap:Body>
// </soap:Envelope>