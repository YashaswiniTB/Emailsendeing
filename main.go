package main

import (
	"bytes"
	"fmt"
	//"net/smtp"
	"text/template"

	//"golang.org/x/tools/go/analysis/passes/nilfunc"
	"gopkg.in/gomail.v2"
)

/*func sendmailsimple(){

   auth :=smtp.PlainAuth(
	"",
	"yashaswinitb0@gmail.com",
	"knsrvuhlwvbkmhwh",
	"smtp.gmail.com",
   )
   msg:="my special subjectbody \n this is body oof mail"
 err :=smtp.SendMail(
	"smtp.gmail.com:587",
	auth,
	"yashaswinitb0@gmail.com",
	[]string{"yashaswinitb0@gmail.com"},
	[]byte(msg),
 )

   if err != nil{
	fmt.Println(err)
 }
}*/
/*func sendmailsimple(subject string, body string, to []string){

	auth :=smtp.PlainAuth(
	 "",
	 "yashaswinitb0@gmail.com",
	 "knsrvuhlwvbkmhwh",
	 "smtp.gmail.com",
	)
	msg := "subject: " + subject + "\n " + body
  err :=smtp.SendMail(
	 "smtp.gmail.com:587",
	 auth,
	 "yashaswinitb0@gmail.com",
	 to,
	 []byte(msg),
  )

	if err != nil{
	 fmt.Println(err)
  }
 }
*/
/*func sendmailsimpleHtml(subject string, html string, to []string){

	auth :=smtp.PlainAuth(
	 "",
	 "yashaswinitb0@gmail.com",
	 "knsrvuhlwvbkmhwh",
	 "smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=:\"UTF-8\";"
	msg := "subject: " + subject + "\n " + headers + "\n\n"+ html
  err :=smtp.SendMail(
	 "smtp.gmail.com:587",
	 auth,
	 "yashaswinitb0@gmail.com",
	 to,
	 []byte(msg),
  )

	if err != nil{
	 fmt.Println(err)
  }
 }*/
/*func sendmailsimpleHtml(subject string, templatePath string, to []string){
    //get ghtml
	var body bytes.Buffer
	t, err  := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string}{ Name: "yashu"} )

	auth := smtp.PlainAuth(
	 "",
	 "yashaswinitb0@gmail.com",
	 "knsrvuhlwvbkmhwh",
	 "smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=:\"UTF-8\";"
	msg := "subject: " + subject + "\n " + headers + "\n\n"+ body.String()
  err =smtp.SendMail(
	 "smtp.gmail.com:587",
	 auth,
	 "yashaswinitb0@gmail.com",
	 to,
	 []byte(msg),
  )

	if err != nil{
	 fmt.Println(err)
  }
 }*/
 func sendgomail(templatepath string){
	var body bytes.Buffer
	t, err:= template.ParseFiles(templatepath)
	t.Execute(&body ,struct { Name string}{Name: "yashu"})
	if err!=nil{
		fmt.Println(err)
		return
	}
	// send withgomail

	m:= gomail.NewMessage()
	m.SetHeader("From","yashaswiniyashuu9@gmail.com")
	m.SetHeader("To","yashaswinitb0@gmail.com")
	m.SetHeader("subject","hello")
	m.SetBody("text/html",/* "hello<b> bob</b> and <i>cora</i>" */body.String())
	m.Attach("./ganesha.gif")
	d:=gomail.NewDialer("smtp.gmail.com",587,"yashaswinitb0@gmail.com","knsrvuhlwvbkmhwh")
   if err:=d.DialAndSend(m); err != nil{
	panic(err)
   }
}
func main(){
	//sendmailsimple("another subject","another b0dy", []string{"yashaswinitb0@gmail.com","yashaswiniyashuu9@gmail.com"})
    /* sendmailsimpleHtml(
		"another subject",
		"<h1> I'm a heading</h1><p> i'm a paragraph</p>",
		[]string{"yashaswinitb0@gmail.com"},
	) */

	/* sendmailsimpleHtml(
		"another subject",
		"./test.html",
		[]string{"yashaswinitb0@gmail.com"},
	)  */
	sendgomail("./test.html")
	
}