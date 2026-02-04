import api from './api';
import type {
  Document,
  DocumentListResponse,
  CreateDocumentRequest,
  UpdateDocumentRequest,
} from './types';

export const documentsApi = {
  // List all documents
  list: async (): Promise<DocumentListResponse> => {
    return api.get<DocumentListResponse>('/docs') as unknown as DocumentListResponse;
  },

  // Get single document
  get: async (name: string): Promise<Document> => {
    return api.get<Document>(`/docs/${name}`) as unknown as Document;
  },

  // Create new document
  create: async (data: CreateDocumentRequest): Promise<Document> => {
    return api.post<Document>('/docs', data) as unknown as Document;
  },

  // Update document
  update: async (name: string, data: UpdateDocumentRequest): Promise<void> => {
    return api.put(`/docs/${name}`, data) as unknown as void;
  },

  // Delete document
  delete: async (name: string): Promise<void> => {
    return api.delete(`/docs/${name}`) as unknown as void;
  },

  // Health check
  health: async (): Promise<{ status: string; service: string; version: string }> => {
    return api.get('/health') as unknown as { status: string; service: string; version: string };
  },
};
