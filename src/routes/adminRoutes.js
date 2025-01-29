const { Router } = require("express");
const {
  addEventCoordinator,
  deleteUser,
} = require("../controllers/adminController.js");
const { authMiddleware } = require("../middlewares/authMiddleware.js");
const { authorizeRole } = require("../middlewares/roleMiddleware.js");
const { param } = require("express-validator");

const router = Router();

router.patch(
  "/add-coordinator/:userId",
  authMiddleware,
  authorizeRole("admin"),
  [param("userId").isUUID().withMessage("User ID must be a UUID")],
  addEventCoordinator
);

router.delete(
  "/delete-user/:userId",
  authMiddleware,
  authorizeRole("admin"),
  [param("userId").isUUID().withMessage("User ID must be a UUID")],
  deleteUser
);

module.exports = router;
