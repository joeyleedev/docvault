import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { documentsApi } from '../services/documents.api';
import type { DocumentMeta, Document } from '../services/types';
import { toast } from 'sonner';

export function useDocuments() {
  return useQuery<DocumentMeta[]>({
    queryKey: ['documents'],
    queryFn: async () => {
      const data = await documentsApi.list();
      return data.items;
    },
    staleTime: 5000, // Consider stale after 5 seconds
  });
}

export function useDocument(name: string) {
  return useQuery<Document>({
    queryKey: ['document', name],
    queryFn: () => documentsApi.get(name),
    enabled: !!name,
    staleTime: 0, // Always refetch when requested
  });
}

export function useCreateDocument() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: documentsApi.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['documents'] });
      toast.success('Document created');
    },
    onError: (error: any) => {
      toast.error(error.message || 'Failed to create document');
    },
  });
}

export function useUpdateDocument() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ name, content }: { name: string; content: string }) =>
      documentsApi.update(name, { content }),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['document', variables.name] });
      queryClient.invalidateQueries({ queryKey: ['documents'] });
    },
  });
}

export function useDeleteDocument() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: documentsApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['documents'] });
      toast.success('Document deleted');
    },
    onError: (error: any) => {
      toast.error(error.message || 'Failed to delete document');
    },
  });
}
