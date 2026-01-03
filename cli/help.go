package cli

var help = `Following are the available commands :- 
	
1. add : to add new tasks
	tags - 
		desc [mandatory] - description of the task being added. 
			Always should be in double commas.
			
		priority - priority of the task being added. 
			Valid values = [low, medium, high, critical]. 

	usage -
		add --desc "testing app"
		add --desc "code review" --priority low

2. list-all : to see all the tasks

3. change : to change the available tasks
	tags -
		id [mandatory] - id of the task being changed.
		
		status - new status of the task. 
			Valid values = [todo, wip, done, cancelled].
		
		priority - new prirority of the task. 
			Valid values = [low, medium, high, critical].
	
	usage -
		change --id 1 --status wip
		change --id 2 --priority critical

3. done : mark any task done
	tags -
		id [mandatory] - id of the task being changed.
	
	usage -
		done --id 1

4. delete : delete any task
	tags -
		id [mandatory] - id of the task being deleted.
	
	usage -
		delete --id 1

5. clear : to clear screen of the terminal
	usage -
		clear
	`
