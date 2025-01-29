const response = require("../response")
const { deleteRecord, getAllRecords, createTable, updateRecord } = require("../utils/sqlFunctions")
const userSchema = require("../schemas/userSchema")

const changeUserRoleToEventCoordinator = async (req, res) => {
  const { userid } = req.body
  try {
    await createTable(userSchema)
    await updateRecord("users", "userid", userid, { role: 'event_coordinator' })
    response(200, "", "User role changed to event coordinator successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const removeUser = async (req, res) => {
  const { userid } = req.params

  try {
    await createTable(userSchema)
    await deleteRecord("users", "userid", userid)
    response(200, "", "User removed successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const getAllUsers = async (req, res) => {
  try {
    await createTable(userSchema)
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
