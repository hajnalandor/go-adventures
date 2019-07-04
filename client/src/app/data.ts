export interface Submit {
  eventType?: string;
  websiteurl?: string;
  time?: number;
}

export interface Dimension {
  width: number;
  height: number;
}

export interface CopyPaste {
  eventType?: string;
  websiteurl: string;
  sessionId: string;
  pasted: boolean;
  formId: string;
}

export interface Resize {
  eventType?: string;
  websiteurl: string;
  sessionId: string;
  resizeFrom?: Dimension;
  resizeTo?: Dimension;

}
