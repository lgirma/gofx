/**
 * Fetches data from a given URL using the Fetch API.
 *
 * @template T
 * @param {string} url - The URL to fetch data from.
 * @param {'GET' | 'POST' | 'PUT' | 'DELETE'} [method='GET'] - The HTTP method to use.
 * @param {unknown} [body=undefined] - The request body for POST, PUT methods.
 * @param {Record<string, string>} [headers={}] - Additional request headers.
 * @returns {Promise<ApiResponse<T>>} A promise that resolves with an ApiResponse object containing data, error, and status.
 */
export async function fetchData(
    url,
    method = 'GET',
    body = undefined,
    headers = {}
) {
    try {
        const options = {
            method,
            ...(body?.constructor === FormData
                ? {}
                : {
                    headers: {
                        'Content-Type': 'application/json',
                        ...(headers ?? {}),
                    },
                }
            ),
        };

        // Include body for methods that typically have one
        if (method !== 'GET' && method !== 'HEAD' && body !== undefined) {
            if (body?.constructor === FormData)
                options.body = body;
            else
                options.body = JSON.stringify(body);
        }

        const response = await fetch(url, options);

        // Fetch API does not throw on HTTP error status codes (like 404, 500).
        // We need to check the response.ok property or status code manually.
        if (!response.ok) {
            let errorMessage = `HTTP error! status: ${response.status}`;
            let errorStatus = response.status;

            // Attempt to parse error message from response body if it's JSON
            try {
                const responseContent = response.headers.get('Content-Type');
                const errorData = responseContent.indexOf('application/json') !== -1
                    ? await response.json()
                    : await response.text();
                errorMessage = errorData.message || errorMessage;
            } catch (e) {
                // If parsing fails, use the default HTTP error message
                console.error("Failed to parse error response body:", e);
            }

            // Handle 401 Unauthorized specifically
            if (errorStatus === 401) {
                // Logout if you have to
            }

            console.error(`API request failed for ${url} with status ${errorStatus}:`, errorMessage);

            return {
                data: null,
                error: errorMessage,
                status: errorStatus,
                headers: response.headers,
                isNetworkError: false,
            };
        }

        const contentType = response.headers.get('Content-Type');
        let data;
        if (contentType.indexOf('application/json') !== -1) {
            data = await response.json();
        } else if (contentType.indexOf('text/') !== -1) {
            data = await response.text();
        } else {
            data = await response.blob();
        }

        return {
            data: data,
            error: null,
            status: response.status,
            headers: response.headers,
        };

    } catch (error) {
        // This catch block primarily handles network errors (e.g., no internet, CORS issues)
        let errorMessage = 'An unexpected error occurred';
        let status = null; // Network errors don't have HTTP status codes

        if (error instanceof Error) {
            errorMessage = error.message;
        }

        console.error(`API request failed for ${url}:`, error);

        return {
            data: null,
            error: errorMessage,
            status,
            isNetworkError: true,
        };
    }
}

/**
 *
 * @param fileName {string}
 * @param blob {Blob}
 * @param contentType {string}
 */
export function download(fileName, blob, contentType) {
    const binaryData = [];
    binaryData.push(blob);
    const blobUrl = URL.createObjectURL(new Blob(binaryData, {type: contentType}));
    const link = document.createElement('a');
    link.href = blobUrl;
    link.download = fileName;
    document.body.appendChild(link);
    link.click();

    setTimeout(() => {
        URL.revokeObjectURL(blobUrl);
        document.body.removeChild(link);
    }, 0);
}

/**
 * @param {ApiResponse} response
 * @param {string} defaultFileName
 * @returns {string}
 */
export function getFileName(response, defaultFileName = "file") {
    const contentDisposition = response.headers.get('content-disposition');
    let filename = defaultFileName;

    if (contentDisposition) {
        const filenameMatch = contentDisposition.match(/filename="(.+)"/);
        if (filenameMatch && filenameMatch.length > 1) {
            filename = filenameMatch[1];
        } else {
            const filenameStarMatch = contentDisposition.match(/filename\*=UTF-8''(.+)$/);
            if (filenameStarMatch && filenameStarMatch.length > 1) {
                try {
                    filename = decodeURIComponent(filenameStarMatch[1]);
                } catch (e) {
                    console.error("Failed to decode filename:", e);
                    filename = filenameStarMatch[1];
                }
            }
        }
    }
    return filename;
}