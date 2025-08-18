/**
 * @returns {AppConfig}
 */
export function getConfig() {
    return {
        apiUrl: import.meta.env.VITE_API_URL,
        apiKey: import.meta.env.VITE_API_KEY,
    };
}