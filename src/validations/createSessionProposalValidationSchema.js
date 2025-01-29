const sessionProposalSchema = {
  userId: {
    in: ["body"],
    isUUID: {
      errorMessage: "User ID must be a UUID",
    },
    notEmpty: {
      errorMessage: "User ID is required",
    },
  },
  title: {
    in: ["body"],
    isString: {
      errorMessage: "Title must be a string",
    },
    notEmpty: {
      errorMessage: "Title is required",
    },
  },
  description: {
    in: ["body"],
    isString: {
      errorMessage: "Description must be a string",
    },
    notEmpty: {
      errorMessage: "Description is required",
    },
  },
  startTime: {
    in: ["body"],
    isISO8601: {
      errorMessage: "Start time must be a valid date",
    },
    notEmpty: {
      errorMessage: "Start time is required",
    },
  },
  endTime: {
    in: ["body"],
    isISO8601: {
      errorMessage: "End time must be a valid date",
    },
    notEmpty: {
      errorMessage: "End time is required",
    },
  },
};

const idSchema = (isDelete = true) => {
  return {
    id: {
      in: ["params"],
      isInt: {
        errorMessage: "Session proposal ID must be an integer",
      },
    },
    userId: {
      in: ["body"],
      isUUID: {
        errorMessage: "User ID must be a UUID",
      },
      optional: isDelete,
    },
  };
};

module.exports = {
  sessionProposalSchema,
  idSchema,
};
