const response = require("../response")


const roleMiddleware = (requiredRole) => {
    return (req, res, next) => {
        const userRole = req.user.role;

        if (userRole !== requiredRole) {
            return response(403, " ", "you are not authorized for this action", res)
        }

        next();
    };
};

module.exports = roleMiddleware;
