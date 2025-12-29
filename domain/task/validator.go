package task

func isValidStatusTransition(currStatus StatusOfTask, newStatus StatusOfTask) bool {
	if currStatus == Todo {
		if newStatus == WIP || newStatus == Cancelled {
			return true
		}
	}

	if currStatus == WIP {
		if newStatus == Done || newStatus == Cancelled {
			return true
		}
	}

	return false
}