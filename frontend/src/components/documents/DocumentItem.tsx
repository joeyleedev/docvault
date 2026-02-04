import { FileText, Trash2 } from 'lucide-react';
import { formatDistanceToNow } from 'date-fns';
import type { DocumentMeta } from '../../services/types';
import { useDeleteDocument } from '../../hooks/useDocuments';
import { useNavigate, useLocation } from 'react-router-dom';
import { useState, useEffect } from 'react';

interface DocumentItemProps {
  document: DocumentMeta;
}

export function DocumentItem({ document }: DocumentItemProps) {
  const deleteDocument = useDeleteDocument();
  const navigate = useNavigate();
  const location = useLocation();
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);

  // Remove .md extension for display
  const displayName = document.name.replace(/\.md$/, '');
  const timeAgo = formatDistanceToNow(new Date(document.updated_at), { addSuffix: true });

  const isActive = location.pathname === `/${displayName}`;

  const handleClick = () => {
    if (!showDeleteConfirm) {
      navigate(`/${displayName}`);
    }
  };

  const handleDelete = async (e: React.MouseEvent) => {
    e.stopPropagation();
    if (showDeleteConfirm) {
      await deleteDocument.mutateAsync(document.name.replace(/\.md$/, ''));
      if (isActive) {
        navigate('/');
      }
      setShowDeleteConfirm(false);
    } else {
      setShowDeleteConfirm(true);
    }
  };

  const handleMouseLeave = () => {
    if (showDeleteConfirm) {
      setShowDeleteConfirm(false);
    }
  };

  // Reset delete confirm state when clicking elsewhere
  useEffect(() => {
    if (!isActive) {
      setShowDeleteConfirm(false);
    }
  }, [isActive]);

  return (
    <li
      className={`
        px-3 py-2 cursor-pointer rounded-md transition-colors duration-150 relative group mx-1
        ${isActive
          ? 'bg-accent text-accent-foreground'
          : 'hover:bg-accent/60 text-foreground'
        }
      `}
      onClick={handleClick}
      onMouseLeave={handleMouseLeave}
    >
      <div className="flex items-center gap-2.5">
        <FileText className="w-4 h-4 flex-shrink-0 text-muted-foreground" />
        <div className="flex-1 min-w-0">
          <p className="text-sm font-medium truncate">{displayName}</p>
          <p className="text-xs text-muted-foreground">{timeAgo}</p>
        </div>
        <button
          className={`
            p-1.5 rounded-md transition-colors duration-150 flex-shrink-0 opacity-0 group-hover:opacity-100
            ${showDeleteConfirm
              ? 'bg-destructive text-destructive-foreground'
              : 'text-muted-foreground hover:bg-destructive/10 hover:text-destructive'
            }
          `}
          onClick={handleDelete}
          title={showDeleteConfirm ? "Confirm delete" : "Delete document"}
        >
          {showDeleteConfirm ? (
            <Trash2 className="w-4 h-4" />
          ) : (
            <Trash2 className="w-4 h-4" />
          )}
        </button>
      </div>
    </li>
  );
}
