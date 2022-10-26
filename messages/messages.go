package messages

import "github.com/fatih/color"

// Databse Messages
var DatabaseNotExist = color.RedString(">> File \"database.json\" does not exist.")
var DatabaseCreated = color.GreenString(">> Creation of \"database.json\" successfully!")
var CreatingDatabase = ">> Creating the \"database.json\" file..."

// Dir Creation (./database)
var NoDirExist = color.RedString(">> Error, the \"./database\" file does not exist!")
var CreateDir = ">> Creation in progress..."
var DirCreated = color.GreenString(">> The \"./database\" file has been successfully created!")

// CreateTask
var IncorrectSyntax = color.RedString(">> Incorrect syntax! Type \"ttm help\" for more information!")
var NewTaskName = ">> Enter the name of the new task: "
var ErrorTaskName = color.RedString(">> Error, you did not enter a name or the name is incorrect. " +
	" Type \"cancel\" to cancel the task creation.")

var NewTaskPriority = ">> Please enter a priority [Low | High]: "
var ErrorTaskPriority = color.RedString(">> Error, you did not enter a priorty or the priorty is incorrect." +
	" Type \"cancel\" to cancel the task creation.")

var DueSoonFeature = color.RedString("Error, feature coming soon, please type \"none\"!")
var DefineDue = ">> How soon do you have to do the task? Type \"none\" to not set a time: "

var ErrorNotInt = color.RedString("Error, the defined time is not a number!")

var SendTaskToDb = ">> Saving in \"database.json\"..."
var TaskCreated = color.GreenString(">> New task created!")

var CancelCreateTask = color.RedString(">> Task creation canceled!")

// Help Command
var HeaderHelpCommand = color.BlueString("##############################\n#         ░H░E░L░P░          #\n#  Ⓣⓔⓡⓜⓘⓝⓐⓛ Ⓣⓐⓢⓚ Ⓜⓐⓝⓐⓖⓜⓔⓝⓣ   #\n##############################")
var CommandsList = ">> ttm create (Create a task)" +
	"\n>> ttm help (See the list of commands)" +
	"\n>> ttm list (See all your tasks)" +
	"\n>> ttm edit (Edit a task)*" +
	"\n>> ttm remove (Remove a task)*" +
	"\n>> ttm clearall (Delete the \"./database\" file)" +
	"\n\n* = Soon"

// List Command
var NoDirOrError = color.RedString(">> Error, you have no task or the \"./database\" file does not exist.")
var NoFileOrError = color.RedString(">> Error, you have no task or the \"./database/database.json\" file does not exist.")

// ClearAll Command
var NoDirOrNoPermission = color.RedString("Error, file does not exist or \"Terminal Task Management\" does not have permissions.")
var DeletingDir = "Deleting the \"./database/\" file..."
var DirDeleted = color.GreenString("Delete \"./database/\" file successfully!")
