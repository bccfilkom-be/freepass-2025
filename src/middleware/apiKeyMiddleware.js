require('dotenv').config();
const API_KEY = process.env.API_KEY;

const apiKeyMiddleware = (req, res, next) => {
    const providedKey = req.headers['x-api-key'];

    if (!providedKey || providedKey !== API_KEY) {
        return res.status(401).json({ 
            error: true,
            code: 401,
            message: 'Unauthorized',
            description: "Permintaan tidak diterapkan karena kurangnya kredensial otentikasi yang valid!"
        });
    }

    next();
}

module.exports = apiKeyMiddleware