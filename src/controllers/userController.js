const { User } = require("../models");
const { validationResult, matchedData } = require("express-validator");
const bcrypt = require("bcrypt");

exports.editProfile = async (req, res) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }

  const {
    params: { id },
  } = req;

  const { username, email, password } = matchedData(req);

  try {
    const user = await User.findByPk(id);
    if (!user) {
      return res.status(404).json({ error: "User not found" });
    }

    if (req.user.id !== id) {
      return res
        .status(403)
        .json({ error: "You are not authorized to perform this action" });
    }

    if (username) user.username = username;
    if (email) user.email = email;
    if (password) user.password = await bcrypt.hash(password, 10);

    await user.save();
    res.status(200).json({ message: "Profile updated successfully", user });
  } catch {
    res.status(500).json({ error: "Internal server error" });
  }
};

exports.viewProfile = async (req, res) => {
  try {
    const userId = req.params.id;

    const user = await User.findOne({
      where: { id: userId },
      attributes: { exclude: ["password"] },
    });

    if (!user) {
      return res.status(404).json({ message: "User not found" });
    }

    res.status(200).json({ user });
  } catch {
    res.status(500).json({ error: "Error retrieving user profile" });
  }
};
