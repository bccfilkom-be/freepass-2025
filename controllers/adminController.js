const response = require("../response")
const { updateRecord, deleteRecord, getAllRecords } = require("../utils/sqlFunctions")

const changeUserRoleToEventCoordinator = async (req, res) => {
  const { userid } = req.body

  try {
    await updateRecord("users", "userid", userid, { role: 'event_coordinator' })
    response(200, "", "User role changed to event coordinator successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const removeUser = async (req, res) => {
  const { userid } = req.params

  try {
    await deleteRecord("users", "userid", userid)
    response(200, "", "User removed successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const getAllUsers = async (req, res) => {
  try {
    const users = await getAllRecords("users")
    response(200, users, "Users retrieved successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

module.exports = {
  changeUserRoleToEventCoordinator,
  removeUser,
  getAllUsers,
}
