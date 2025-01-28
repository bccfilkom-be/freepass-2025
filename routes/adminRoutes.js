const express = require("express")
const { addEventCoordinator, removeUser, getAllUsers } = require("../controllers/adminController")
const authMiddleware = require("../middleware/authMiddleware")
const roleMiddleware = require("../middleware/roleMiddleware")
const router = express.Router()

router.post("/coordinator", authMiddleware, roleMiddleware("admin"), addEventCoordinator)
router.delete("/user/:userid", authMiddleware, roleMiddleware("admin"), removeUser)
router.get("/users", authMiddleware, roleMiddleware("admin"), getAllUsers)

module.exports = router
