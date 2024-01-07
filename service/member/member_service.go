package member

import (
	constants "github.com/infinity-api/constant"
	"github.com/infinity-api/dto"
	"github.com/infinity-api/model"
	repo "github.com/infinity-api/repository"
)

type MemberService interface {
	GetAllMembers() ([]model.Member, error)
}

type memberService struct {
	memberRepository repo.MemberRepository
}

func NewMemberService(memberRepository repo.MemberRepository) MemberService {
	return memberService{memberRepository: memberRepository}
}

func (service memberService) GetAllMembers() ([]model.Member, error) {
	memberDto, err := service.memberRepository.GetMembers()
	if err != nil {
		return nil, err
	}
	members := buildMembers(memberDto)
	return members, err
}

func buildMembers(membersDto []dto.MemberDTO) []model.Member {
	var members []model.Member
	for _, memberDTo := range membersDto {
		member := buildMember(memberDTo)
		members = append(members, member)
	}
	return members
}

func buildMember(memberDTo dto.MemberDTO) model.Member {
	member := model.Member{
		Id:              *memberDTo.Id,
		FamilyNo:        *memberDTo.FamilyNo,
		SubscriptionNo:  *memberDTo.SubscriptionNo,
		FirstName:       memberDTo.FirstName,
		LastName:        memberDTo.LastName,
		NameEng:         memberDTo.EngLastName,
		BaptisedNameEng: memberDTo.EngBaptisedName,
		Gender:          constants.Gender(*memberDTo.Gender),
		MaritalStatus:   constants.MaritalStatus(*memberDTo.MaritalStatus),
		Address: model.Address{
			DoorNo:   0,
			Street:   memberDTo.Street,
			Town:     memberDTo.Town,
			District: memberDTo.City,
			PinCode:  memberDTo.PostalCode,
		},
		FamilyMembers: nil,
	}
	member.Title = member.GetTitle()
	member.FullName = member.GetFullName()
	return member
}
