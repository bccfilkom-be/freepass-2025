const { User } = require("../models");
const { validationResult, matchedData } = require("express-validator");

exports.addEventCoordinator = async (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }

  const { userId } = matchedData(req);

  try {
    const user = await User.findByPk(userId);

    if (!user) {
      return res.status(404).json({ message: "User not found" });
    }

    if (user.role === "coordinator") {
      return res
        .status(400)
        .json({ message: "User is already an event coordinator" });
    }

    await user.update({
      role: "coordinator",
    });

    res.status(200).json({ message: "User is now an event coordinator", user });
  } catch {
    res.status(500).json({ message: "Failed to update user role" });
  }
};

exports.deleteUser = async (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }

  const { userId } = matchedData(req);
  const {
    user: { role },
  } = req;

  try {
    const user = await User.findByPk(userId);

    if (!user) {
      return res.status(404).json({ message: "User not found" });
    }

    if (user.role === "admin" && role === "admin") {
      return res.status(400).json({ message: "Cannot delete other admins" });
    }

    await user.destroy();

    res.status(200).json({ message: "User deleted successfully" });
  } catch {
    res.status(500).json({ message: "Failed to delete user" });
  }
};
