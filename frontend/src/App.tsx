import { createBrowserRouter, RouterProvider, useParams, Navigate } from 'react-router-dom';
import { Layout } from './components/layout/Layout';
import { EmptyState } from './components/documents/EmptyState';
import { DocumentEditor } from './components/editor/DocumentEditor';
import { useDocument } from './hooks/useDocuments';

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        index: true,
        element: <EmptyState />,
      },
      {
        path: ':name',
        element: <DocumentRoute />,
      },
    ],
  },
]);

function DocumentRoute() {
  const { name } = useParams();
  const { data: document, isLoading, error } = useDocument(name || '');

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-muted-foreground">Loading...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-destructive">Document not found</div>
      </div>
    );
  }

  if (!document) {
    return <Navigate to="/" replace />;
  }

  return <DocumentEditor documentName={name || ''} initialContent={document.content} />;
}

export default function App() {
  return <RouterProvider router={router} />;
}
