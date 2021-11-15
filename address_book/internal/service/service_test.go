package service

import (
	"addressbook/internal/mock"
	models "addressbook/internal/model"
	"addressbook/internal/pb"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

func TestService(t *testing.T) {
	suite.Run(t, new(Suite))
}

type Suite struct {
	suite.Suite
	db mock.GormInterface
	a  AddressBookService
}

func (s *Suite) SetupTest() {
	s.db = mock.GormInterface{}
	s.a = AddressBookService{
		db: &s.db,
	}
}

func (s *Suite) TestMajorAddContact() {
	ctx := context.Background()
	tx := &gorm.DB{}

	contact := &models.Contact{
		Name:    "alex",
		Phone:   "89871111111",
		Address: "my house",
	}

	resp := &pb.AddContactRequest{
		Contact: &pb.Contact{
			Name:    "alex",
			Phone:   "89871111111",
			Address: "my house",
		},
	}

	s.db.On("FirstOrCreate", contact).Once().Return(tx)
	req, err := s.a.AddContact(ctx, resp)
	s.NoError(err)
	s.Equal(successAdding, req.Msg)
}

func (s *Suite) TestMinorAddContact() {
	ctx := context.Background()

	emptyResp := &pb.AddContactRequest{}

	req, err := s.a.AddContact(ctx, nil)
	s.NoError(err)
	s.Equal(invalidInputData, req.Msg)

	req, err = s.a.AddContact(ctx, emptyResp)
	s.NoError(err)
	s.Equal(invalidInputData, req.Msg)

	resp := &pb.AddContactRequest{
		Contact: &pb.Contact{
			Name:    "alex",
			Phone:   "89871111111",
			Address: "my house",
		},
	}

	tx := &gorm.DB{
		Error: errors.New("some error"),
	}

	contact := &models.Contact{
		Name:    "alex",
		Phone:   "89871111111",
		Address: "my house",
	}
	s.db.On("FirstOrCreate", contact).Once().Return(tx)

	req, err = s.a.AddContact(ctx, resp)
	s.Error(err)
	s.Equal(fmt.Sprintf("%s, %v", addError, "some error"), req.Msg)
}

func (s *Suite) TestMajorDeleteContact() {
	ctx := context.Background()
	tx := &gorm.DB{}

	contact := &models.Contact{
		Name:    "alex",
		Phone:   "89871111111",
		Address: "my house",
	}

	resp := &pb.DeleteContactRequest{
		Phone: "89871111111",
	}

	s.db.On("Delete", &models.Contact{}, contact.Phone).Once().Return(tx)
	req, err := s.a.DeleteContact(ctx, resp)
	s.NoError(err)
	s.Equal(successDeleting, req.Msg)
}

func (s *Suite) TestMinorDeleteContact() {
	ctx := context.Background()

	req, err := s.a.DeleteContact(ctx, nil)
	s.NoError(err)
	s.Equal(invalidInputData, req.Msg)

	tx := &gorm.DB{
		Error: errors.New("some error"),
	}

	contact := &models.Contact{
		Name:    "alex",
		Phone:   "89871111111",
		Address: "my house",
	}

	resp := &pb.DeleteContactRequest{
		Phone: "89871111111",
	}

	s.db.On("Delete", &models.Contact{}, contact.Phone).Once().Return(tx)
	req, err = s.a.DeleteContact(ctx, resp)
	s.Error(err)
	s.Equal(fmt.Sprintf("%s, %v", deleteError, "some error"), req.Msg)
}

type StructImplDialectorInterface struct{}

func (a *StructImplDialectorInterface) Name() string {
	return ""
}

func (a *StructImplDialectorInterface) Initialize(*gorm.DB) error {
	return nil
}

func (a *StructImplDialectorInterface) Migrator(db *gorm.DB) gorm.Migrator {
	return nil
}

func (a *StructImplDialectorInterface) DataTypeOf(*schema.Field) string {
	return ""
}

func (a *StructImplDialectorInterface) DefaultValueOf(*schema.Field) clause.Expression {
	return nil
}

func (a *StructImplDialectorInterface) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
}

func (a *StructImplDialectorInterface) QuoteTo(clause.Writer, string) {}

func (a *StructImplDialectorInterface) Explain(sql string, vars ...interface{}) string {
	return ""
}

func (s *Suite) TestMajorUpdateContact() {
	ctx := context.Background()
	d := StructImplDialectorInterface{}
	tx, _ := gorm.Open(&d, &gorm.Config{})

	contact := &models.Contact{
		Name:    "alex",
		Phone:   "89871111111",
		Address: "my house",
	}

	resp := &pb.UpdateContactRequest{
		Contact: &pb.Contact{
			Name:    contact.Name,
			Phone:   contact.Phone,
			Address: contact.Address,
		},
	}
	s.db.On("Model", contact).Return(tx).On("Updates", models.Contact{}).Once().Return(tx)
	req, err := s.a.UpdateContact(ctx, resp)
	s.NoError(err)
	s.Equal(successUpdating, req.Msg)
}

func (s *Suite) TestMinorUpdateContact() {
	ctx := context.Background()
	emptyResp := &pb.UpdateContactRequest{}

	req, err := s.a.UpdateContact(ctx, nil)
	s.NoError(err)
	s.Equal(invalidInputData, req.Msg)

	req, err = s.a.UpdateContact(ctx, emptyResp)
	s.NoError(err)
	s.Equal(invalidInputData, req.Msg)

	d := StructImplDialectorInterface{}
	tx, _ := gorm.Open(&d, &gorm.Config{})
	tx.Error = errors.New("some error")

	contact := &models.Contact{
		Name:    "alex",
		Phone:   "89871111111",
		Address: "my house",
	}

	resp := &pb.UpdateContactRequest{
		Contact: &pb.Contact{
			Name:    contact.Name,
			Phone:   contact.Phone,
			Address: contact.Address,
		},
	}
	s.db.On("Model", contact).Return(tx).On("Updates", models.Contact{}).Once().Return(tx)
	req, err = s.a.UpdateContact(ctx, resp)
	s.Error(err)
	s.Equal(fmt.Sprintf("%s, %v", updateError, "some error"), req.Msg)
}

func (s *Suite) TestMajor() {
	ctx := context.Background()
	d := StructImplDialectorInterface{}
	tx, _ := gorm.Open(&d, &gorm.Config{})

	resp := &pb.FindContactRequest{
		Query:      "Alex",
		SearchType: 0,
	}

	contacts := []models.Contact{}

	s.db.On("Where", "name = ?", resp.Query).Return(tx).On("Find", []models.Contact{}).Once().Return(tx)
	req, err := s.a.FindContact(ctx, resp)
	s.NoError(err)
	s.Equal(getFindContactMsg(len(contacts)), req.Msg)

	resp = &pb.FindContactRequest{
		Query:      "89871111111",
		SearchType: 1,
	}

	s.db.On("Where", "phone LIKE ?", resp.Query).Return(tx).On("Find", []models.Contact{}).Once().Return(tx)
	req, err = s.a.FindContact(ctx, resp)
	s.NoError(err)
	s.Equal(getFindContactMsg(len(contacts)), req.Msg)

	resp = &pb.FindContactRequest{
		Query:      "89871111111",
		SearchType: 2,
	}

	req, err = s.a.FindContact(ctx, resp)
	s.NoError(err)
	s.Equal(wrongValueSearchError, req.Msg)
}

func (s *Suite) TestMinor() {
	ctx := context.Background()
	emptyResp := &pb.FindContactRequest{}

	req, err := s.a.FindContact(ctx, nil)
	s.NoError(err)
	s.Equal(invalidInputData, req.Msg)

	req, err = s.a.FindContact(ctx, emptyResp)
	s.NoError(err)
	s.Equal(invalidInputData, req.Msg)
}

func (s *Suite) TestProcessFindContact() {
	findContacts := []models.Contact{}

	req, err := processFindContact(&findContacts, nil)
	s.NoError(err)
	s.Equal(getFindContactMsg(len(findContacts)), req.Msg)

	req, err = processFindContact(&findContacts, errors.New("some error"))
	s.Error(err)
	s.Equal(fmt.Sprintf("%s, %v", findError, "some error"), req.Msg)
}

func (s *Suite) TestGetFindContactMsg() {
	s.Equal(getFindContactMsg(0), findZeroContactMsg)
	s.Equal(getFindContactMsg(1), findOneContactMsg)
	s.Equal(getFindContactMsg(2), fmt.Sprintf("%s %v", findSomeContactsMsg, 2))
}

func (s *Suite) TestNewAddressBookService() {
	db := &gorm.DB{}
	service := NewAddressBookService(db)
	s.Equal(service.db, db)
}
