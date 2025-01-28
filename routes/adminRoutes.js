const express = require("express")
const { changeUserRoleToEventCoordinator, removeUser } = require("../controllers/adminController")
const authMiddleware = require("../middleware/authMiddleware")
const roleMiddleware = require("../middleware/roleMiddleware")
const router = express.Router()

router.put("/coordinator", authMiddleware, roleMiddleware("admin"), changeUserRoleToEventCoordinator)
router.delete("/user/:userid", authMiddleware, roleMiddleware("admin"), removeUser)

module.exports = router
