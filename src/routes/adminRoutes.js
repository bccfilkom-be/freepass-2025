const { Router } = require("express");
const {
  addEventCoordinator,
  deleteUser,
} = require("../controllers/adminController.js");
const { authMiddleware } = require("../middlewares/authMiddleware.js");
const { authorizeRole } = require("../middlewares/roleMiddleware.js");

const router = Router();

router.patch(
  "/add-coordinator/:userId",
  authMiddleware,
  authorizeRole("admin"),
  addEventCoordinator
);
router.delete(
  "/delete-user/:userId",
  authMiddleware,
  authorizeRole("admin"),
  deleteUser
);

module.exports = router;
