import { useEditor, EditorContent } from '@tiptap/react';
import StarterKit from '@tiptap/starter-kit';
import Placeholder from '@tiptap/extension-placeholder';
import { useUpdateDocument } from '../../hooks/useDocuments';
import { useEffect } from 'react';
import styles from './Editor.module.css';

interface DocumentEditorProps {
  documentName: string;
  initialContent: string;
}

export function DocumentEditor({ documentName, initialContent }: DocumentEditorProps) {
  const updateDocument = useUpdateDocument();

  const editor = useEditor({
    extensions: [
      StarterKit.configure({
        heading: { levels: [1, 2, 3] },
        codeBlock: false,
        horizontalRule: false,
      }),
      Placeholder.configure({
        placeholder: 'Start writing...',
      }),
    ],
    content: initialContent,
    editorProps: {
      attributes: {
        class: styles.editor,
      },
    },
    onUpdate: ({ editor }) => {
      const htmlContent = editor.getHTML();
      updateDocument.mutate({ name: documentName, content: htmlContent });
    },
  });

  // Handle external content updates (e.g., when switching documents)
  useEffect(() => {
    if (editor && editor.getHTML() !== initialContent) {
      editor.commands.setContent(initialContent, { emitUpdate: false });
    }
  }, [initialContent, editor]);

  if (!editor) {
    return null;
  }

  return (
    <div className={styles.container}>
      <EditorContent editor={editor} />
    </div>
  );
}
