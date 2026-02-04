import { useState } from 'react';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from '../ui/dialog';
import { Input } from '../ui/input';
import { Button } from '../ui/button';
import { useCreateDocument } from '../../hooks/useDocuments';
import { useNavigate } from 'react-router-dom';

interface CreateDialogProps {
  isOpen: boolean;
  onClose: () => void;
}

export function CreateDialog({ isOpen, onClose }: CreateDialogProps) {
  const [name, setName] = useState('');
  const createDocument = useCreateDocument();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!name.trim()) {
      return;
    }

    try {
      const cleanName = name.replace(/\.md$/, '');
      await createDocument.mutateAsync({
        name: cleanName,
        content: '',
      });

      setName('');
      onClose();
      navigate(`/${cleanName}`);
    } catch (error) {
      // Error is handled by the hook
    }
  };

  const handleClose = () => {
    setName('');
    onClose();
  };

  return (
    <Dialog open={isOpen} onOpenChange={handleClose}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>New Document</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <Input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="my-document"
              autoFocus
            />
          </div>

          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={handleClose}
              disabled={createDocument.isPending}
            >
              Cancel
            </Button>
            <Button
              type="submit"
              disabled={createDocument.isPending}
            >
              Create
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
