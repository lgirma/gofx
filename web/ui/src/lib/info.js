import {getConfig} from "./config.js";
import {fetchData} from "$lib/fetch.js";

/**
 * @returns {Promise<ApiResponse<ApiInfo>>}
 */
export async function getAppInfo() {
    const {apiUrl, apiKey} = getConfig();
    return await fetchData(
        `${apiUrl}/status/ping?apiKey=${apiKey}`,
        "GET",
    );
}