package main

import (
	"fmt"
)

type Client struct {
	clientId string
	name     string
	surname  string
}

func newClient(clientId string, name string, surname string) *Client {
	return &Client{
		clientId: clientId,
		name:     name,
		surname:  surname,
	}
}

type Facade struct {
	client       *Client
	account      *account
	securityCode *securityCode
	//notification *notification
	//ledger       *ledger
}

func newClientFacade(name string, surname string, accountID string) *Facade {
	fmt.Println("Starting create account")
	wallet := &Facade{
		client:       newClient(accountID, name, surname),
		account:      newAccount(accountID),
		securityCode: newSecurityCode(1234),
		//wallet:       newWallet(sum),
		//notification: &notification{},
		//ledger: &ledger{},
	}
	fmt.Println("Account created")
	return wallet
}

type account struct {
	accountID string
}

func newAccount(accountID string) *account {
	return &account{
		accountID: accountID,
	}
}

func (a *account) checkAccount(accountID string) error {
	if a.accountID != accountID {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

func buyTicket() error {
	fmt.Println("Purchase made. Thank you for choosing our travel agency.")
	return nil
}

func menuview() {
	fmt.Println("Create account")
	var text string
	fmt.Print("Select: ")
	fmt.Scanf("%s", &text)
	selectMenu(text)
}

func selectMenu(choose string) {
	if choose == "create" {
		var name string
		var surname string
		var id string

		fmt.Scanf("%s\n ", &name)
		fmt.Scanf("%s\n", &surname)
		fmt.Scanf("%s\n", &id)

		newaccount := newClientFacade(name, surname, id)

		fmt.Println(newaccount.client.name)
	} else {
		fmt.Println("Выберите правильно")
		menuview()
	}

}

func main() {
	menuview()

}
