const { Router } = require("express");
const authRoutes = require("./authRoutes.js");
const userRoutes = require("./userRoutes.js");

const router = Router();

router.use("/auth", authRoutes);
router.use("/users", userRoutes);

module.exports = router;
