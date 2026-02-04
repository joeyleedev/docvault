// API Response Types
export interface ApiSuccessResponse<T> {
  code: string;
  message: string;
  data: T;
}

export interface ApiErrorResponse {
  code: string;
  message: string;
  details?: string;
}

// Document Types
export interface DocumentMeta {
  name: string;           // includes .md extension
  updated_at: string;      // ISO 8601 date
  size: number;
}

export interface Document {
  name: string;
  content: string;
  created_at?: string;
  updated_at?: string;
  size?: number;
}

export interface DocumentListResponse {
  total: number;
  items: DocumentMeta[];
}

// Request Types
export interface CreateDocumentRequest {
  name: string;
  content: string;
}

export interface UpdateDocumentRequest {
  content: string;
}
