package dbconnection

func buildForUpdate(temp Member) []string {
	temp1 := []string{}
	if temp.Gender != "" {
		temp1 = append(temp1, "GENDER")
	}
	if temp.SkinColor != "" {
		temp1 = append(temp1, "SKINCOLOR")
	}
	if temp.SkinType != "" {
		temp1 = append(temp1, "SKINTYPE")
	}
	if temp.UserName != "" {
		temp1 = append(temp1, "USERNAME")
	}
	return temp1
}
