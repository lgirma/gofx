interface ApiResponse<T = any> {
    status: number;
    data?: T;
    error?: string;
    headers?: Headers;
    isNetworkError: boolean;
}

interface AppConfig {
    apiUrl: string;
}

interface ApiInfo {
    env: string;
    version: string;
}