package git

import (
	"data"
	"os"
)

func isExist(repository data.Repository) bool{
	if _, err := os.Stat("./"+repository.Owner.Name + "/" + repository.Name); os.IsNotExist(err) {
		return false
	}else {
		return true
	}
}

func Update(repository data.Repository) bool{
	//Check if repository is already tracked
	if isExist(repository){
		pull(repository)
	}else{
		clone(repository)
	}
	return false
}

func clone(repository data.Repository) bool{
	println("Cloning...")
	return false
}

func pull(repository data.Repository) bool{
	println("Pulling...")
	return false
}