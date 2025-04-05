export interface PaginatedResponse<T> {
    data: T[];
    total: number;
    limit: number;
    page: number;
    totalPages: number;
}