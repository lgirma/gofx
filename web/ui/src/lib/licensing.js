import {fetchData} from "$lib/fetch.js";
import {getConfig} from "./config.js";

/**
 *
 * @returns {Promise<ApiResponse<boolean>>}
 */
export async function isLicensed() {
    const {apiUrl} = getConfig();
    return await fetchData(
        `${apiUrl}/license/is-installed`,
        "GET",
    );
}

/**
 *
 * @returns {Promise<ApiResponse<string>>}
 */
export async function getRequestCode() {
    const {apiUrl} = getConfig();
    return await fetchData(
        `${apiUrl}/license/request-code`,
        "GET",
    );
}

/**
 * @param {string} activationCode
 * @returns {Promise<ApiResponse<string>>}
 */
export async function activate(activationCode) {
    const {apiUrl} = getConfig();
    return await fetchData(
        `${apiUrl}/license/activate`,
        "POST",
        {activationCode},
    );
}