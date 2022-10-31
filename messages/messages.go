package messages

import "github.com/fatih/color"

// Clear Terminal Function
var ErrorWhenClearTerminal = color.RedString(">> Error while cleaning the terminal. Please contact Kstars#5038 via Discord!")

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

var DueSoonFeature = color.RedString(">> Error, feature coming soon, please type \"none\"!")
var DefineDue = ">> How soon do you have to do the task? Type \"none\" to not set a time: "

var ErrorNotIntDueTime = color.RedString(">> Error, the defined time is not a number!")

var SendTaskToDb = ">> Saving in \"database.json\"..."
var TaskCreated = color.GreenString(">> New task created!")

var CancelCreateTask = color.RedString(">> Task creation canceled!")

// Help Command
var HeaderHelpCommand = color.BlueString("##############################\n#         ░H░E░L░P░          #\n#  Ⓣⓔⓡⓜⓘⓝⓐⓛ Ⓣⓐⓢⓚ ⓜⓐⓝⓐⓖⓜⓔⓝⓣ   #\n##############################")
var CommandsList = ">> ttm create (Create a task)" +
	"\n>> ttm help (See the list of commands)" +
	"\n>> ttm list (See all your edit_tasks)" + //METTRE COULEUR AU STATUS & PRIORITY
	"\n>> ttm edit (Edit a task)" + color.YellowString("*") + //A FINIR
	"\n>> ttm remove (Remove a task)" +
	"\n>> ttm options (Change Terminal Task Management options)" + color.RedString("*") + //A FAIRE
	"\n>> ttm clearall (Delete the \"./database\" file)" +
	"\n\n" + color.RedString("*") + " = Feature to do | " + color.YellowString("*") + " = Feature in development"

// List Command
var NoDirOrError = color.RedString(">> Error, you have no task or the \"./database\" file does not exist.")
var NoFileOrError = color.RedString(">> Error, you have no task or the \"./database/database.json\" file does not exist.")

// ClearAll Command
var NoDirOrNoPermission = color.RedString(">> Error, file does not exist or \"Terminal Task Management\" does not have permissions.")
var DeletingDir = ">> Deleting the \"./database/\" file..."
var DirDeleted = color.GreenString(">> Delete \"./database/\" file successfully!")
var ConfirmClearAll = ">> Are you sure you want to delete the " + color.RedString("\"./database/*\" ") + "file? [Y/N]: "
var NotGoodResponse = color.RedString(">> I didn't understand your answer!")
var ClearAllCanceled = color.GreenString(">> You have canceled the deletion of the \"./database/*\" file!")

// Remove Task Command
var DefineIDTask = ">> Please enter task ID, Type \"cancel\" to cancel the action: "
var ErrorNotIntTask = color.RedString(">> Error, the defined ID is not a number!")
var TaskNotFoundSearhTask = color.RedString(">> Error, task does not exist.")
var RemoveTaskCanceled = color.RedString(">> Task deletion has been cancelled!")

// Options Command
var CommandOptionsSoon = color.RedString(">> Error, this feature is not yet available!")

// Edit Command
var HeaderEditCommand = color.BlueString("##############################\n#         ░E░D░I░T░          #\n#  Ⓣⓔⓡⓜⓘⓝⓐⓛ Ⓣⓐⓢⓚ ⓜⓐⓝⓐⓖⓜⓔⓝⓣ   #\n##############################")
var ChooseOneEdit = "[1] Edit task name" + color.YellowString("*")
var ChooseTwoEdit = "[2] Change task status"
var ChooseThreeEdit = "[3] Change task priority" + color.RedString("*")
var ChooseForEdit = "[4] Change task due" + color.RedString("*")
var ChooseFiveEdit = "[5] Close edit menu" + "\n\n" + color.RedString("*") + " = Feature to do | " + color.YellowString("*") + " = Feature in development"

var EditMenuClosed = color.RedString(">> You have closed the edit options menu")

var DefineEditChoose = ">> Enter the number of an option to modify: "
var CommandEditSoon = color.RedString(">> Error, this feature is not yet available!")

// ## EDIT NAME TASK ##//
var CancelEditTask = color.RedString(">> You canceled the task modification.")
var EnterNewNameTask = ">> Please enter the new name to set to the task: "
var TaskNameEdited = color.GreenString(">> Your task has been modified successfully!")

// ## EDIT STATUS TASK ##//
var ConfirmDoneTask = ">> Are you sure you want to switch the task to \"Done\" status? [Y/N]: "
var TaskEdited = color.GreenString(">> The task modification has been successfully completed!")

// ## EDIT PRIORITY TASK ##//
var DefineNewPriority = ">> Please set the new priority [Low | High]: "
