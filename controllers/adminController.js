const response = require("../response")
const { insertRecord, deleteRecord, getAllRecords, createTable } = require("../utils/sqlFunctions")
const userSchema = require("../schemas/userSchema")

const addEventCoordinator = async (req, res) => {
  const coordinatorData = req.body

  try {
    await createTable(userSchema)
    const newCoordinator = await insertRecord("users", coordinatorData)
    response(201, newCoordinator, "Event coordinator added successfully", res)
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
  addEventCoordinator,
  removeUser,
  getAllUsers,
}
