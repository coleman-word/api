package design
 
import (
 . "github.com/goadesign/goa/design"
 . "github.com/goadesign/goa/design/apidsl"
)
 
var _ = API("Word", func() {
 Title("Word Services")
 Description("Create your own Dictionary!")
 Version("1.0")
 BasePath("/word")
 Scheme("http")
 Host("localhost:8080")
 Consumes("application/json")
})
