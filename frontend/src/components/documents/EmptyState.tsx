import { FileText } from 'lucide-react';

export function EmptyState() {
  return (
    <div className="flex flex-col items-center justify-center h-full text-center p-8">
      <div className="flex flex-col items-center gap-4 max-w-md">
        <div className="p-4 rounded-full bg-muted">
          <FileText className="w-12 h-12 text-muted-foreground" />
        </div>
        <div>
          <h2 className="text-xl font-semibold mb-2">No document selected</h2>
          <p className="text-muted-foreground mb-4">
            Select a document from the sidebar or create a new one to get started.
          </p>
        </div>
      </div>
    </div>
  );
}
