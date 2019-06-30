export interface Data {
  websiteurl?: string;
  resizeFrom?: Dimension;
  resizeTo?: Dimension;
  copyAndPaste?: boolean;
  formCompletionTime?: number;
}

export interface Dimension {
  width: number;
  height: number;
}
