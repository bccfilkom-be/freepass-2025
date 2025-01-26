const { Router } = require("express");
const authRoutes = require("./authRoutes.js");
const userRoutes = require("./userRoutes.js");
const sessionRoutes = require("./sessionRoutes.js");

const router = Router();

router.use("/auth", authRoutes);
router.use("/users", userRoutes);
router.use("/sessions", sessionRoutes);

module.exports = router;
