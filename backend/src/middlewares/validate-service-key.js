const { ApiError } = require("../utils");
const { env } = require("../config");

const validateServiceKey = (req, res, next) => {
    const serviceKey = req.headers["x-service-key"];

    if (!serviceKey || serviceKey !== env.SERVICE_API_KEY) {
        throw new ApiError(401, "Invalid or missing service key");
    }

    next();
};

module.exports = { validateServiceKey };
