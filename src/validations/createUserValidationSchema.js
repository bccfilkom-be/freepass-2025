const { User } = require("../models");

const createUserValidationSchema = (isUpdate = false, allowRole = false) => {
  return {
    username: {
      in: ["body"],
      isEmpty: {
        negated: true,
        errorMessage: "Username is required",
      },
      isString: {
        errorMessage: "Username must be a string",
      },
      optional: isUpdate,
    },
    email: {
      in: ["body"],
      isEmail: {
        errorMessage: "Invalid email format",
      },
      custom: {
        options: async (email) => {
          const existingUser = await User.findOne({ where: { email } });
          if (existingUser) {
            throw new Error("Email already in use");
          }
        },
      },
      optional: isUpdate,
    },
    password: {
      in: ["body"],
      isLength: {
        options: { min: 8 },
        errorMessage: "Password must be at least 8 characters long",
      },
      optional: isUpdate,
    },
    ...(allowRole && {
      role: {
        in: ["body"],
        isIn: {
          options: [["user", "admin"]],
          errorMessage: "Invalid role provided",
        },
        optional: true,
      },
    }),
  };
};

module.exports = {
  createUserValidationSchema,
};
