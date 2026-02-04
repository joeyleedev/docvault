import { useDocuments } from '../../hooks/useDocuments';
import { DocumentItem } from './DocumentItem';
import { ScrollArea } from '../ui/scroll-area';
import { Loader2 } from 'lucide-react';

export function DocumentList() {
  const { data: documents, isLoading, error } = useDocuments();

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full p-4">
        <Loader2 className="w-5 h-5 animate-spin text-muted-foreground" />
      </div>
    );
  }

  if (error) {
    return (
      <div className="p-4 text-sm text-destructive">
        Failed to load documents
      </div>
    );
  }

  if (!documents || documents.length === 0) {
    return (
      <div className="p-4 text-sm text-muted-foreground text-center">
        No documents yet
      </div>
    );
  }

  return (
    <ScrollArea className="flex-1">
      <ul className="p-2">
        {documents.map((doc) => (
          <DocumentItem
            key={doc.name}
            document={doc}
          />
        ))}
      </ul>
    </ScrollArea>
  );
}
