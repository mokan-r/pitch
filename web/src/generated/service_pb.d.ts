import * as jspb from 'google-protobuf'



export class Music extends jspb.Message {
  getSong(): Song | undefined;
  setSong(value?: Song): Music;
  hasSong(): boolean;
  clearSong(): Music;

  getElapsed(): number;
  setElapsed(value: number): Music;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Music.AsObject;
  static toObject(includeInstance: boolean, msg: Music): Music.AsObject;
  static serializeBinaryToWriter(message: Music, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Music;
  static deserializeBinaryFromReader(message: Music, reader: jspb.BinaryReader): Music;
}

export namespace Music {
  export type AsObject = {
    song?: Song.AsObject,
    elapsed: number,
  }
}

export class Song extends jspb.Message {
  getId(): string;
  setId(value: string): Song;

  getSongtitle(): string;
  setSongtitle(value: string): Song;

  getDuration(): number;
  setDuration(value: number): Song;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Song.AsObject;
  static toObject(includeInstance: boolean, msg: Song): Song.AsObject;
  static serializeBinaryToWriter(message: Song, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Song;
  static deserializeBinaryFromReader(message: Song, reader: jspb.BinaryReader): Song;
}

export namespace Song {
  export type AsObject = {
    id: string,
    songtitle: string,
    duration: number,
  }
}

export class NextRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NextRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NextRequest): NextRequest.AsObject;
  static serializeBinaryToWriter(message: NextRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NextRequest;
  static deserializeBinaryFromReader(message: NextRequest, reader: jspb.BinaryReader): NextRequest;
}

export namespace NextRequest {
  export type AsObject = {
  }
}

export class NextResponse extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): NextResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NextResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NextResponse): NextResponse.AsObject;
  static serializeBinaryToWriter(message: NextResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NextResponse;
  static deserializeBinaryFromReader(message: NextResponse, reader: jspb.BinaryReader): NextResponse;
}

export namespace NextResponse {
  export type AsObject = {
    ok: boolean,
  }
}

export class PrevRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PrevRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PrevRequest): PrevRequest.AsObject;
  static serializeBinaryToWriter(message: PrevRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PrevRequest;
  static deserializeBinaryFromReader(message: PrevRequest, reader: jspb.BinaryReader): PrevRequest;
}

export namespace PrevRequest {
  export type AsObject = {
  }
}

export class PrevResponse extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): PrevResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PrevResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PrevResponse): PrevResponse.AsObject;
  static serializeBinaryToWriter(message: PrevResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PrevResponse;
  static deserializeBinaryFromReader(message: PrevResponse, reader: jspb.BinaryReader): PrevResponse;
}

export namespace PrevResponse {
  export type AsObject = {
    ok: boolean,
  }
}

export class AddSongRequest extends jspb.Message {
  getSongId(): string;
  setSongId(value: string): AddSongRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddSongRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddSongRequest): AddSongRequest.AsObject;
  static serializeBinaryToWriter(message: AddSongRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddSongRequest;
  static deserializeBinaryFromReader(message: AddSongRequest, reader: jspb.BinaryReader): AddSongRequest;
}

export namespace AddSongRequest {
  export type AsObject = {
    songId: string,
  }
}

export class AddSongResponse extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): AddSongResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddSongResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddSongResponse): AddSongResponse.AsObject;
  static serializeBinaryToWriter(message: AddSongResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddSongResponse;
  static deserializeBinaryFromReader(message: AddSongResponse, reader: jspb.BinaryReader): AddSongResponse;
}

export namespace AddSongResponse {
  export type AsObject = {
    ok: boolean,
  }
}

export class CreateSongRequest extends jspb.Message {
  getSong(): Song | undefined;
  setSong(value?: Song): CreateSongRequest;
  hasSong(): boolean;
  clearSong(): CreateSongRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSongRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSongRequest): CreateSongRequest.AsObject;
  static serializeBinaryToWriter(message: CreateSongRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSongRequest;
  static deserializeBinaryFromReader(message: CreateSongRequest, reader: jspb.BinaryReader): CreateSongRequest;
}

export namespace CreateSongRequest {
  export type AsObject = {
    song?: Song.AsObject,
  }
}

export class CreateSongResponse extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): CreateSongResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSongResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSongResponse): CreateSongResponse.AsObject;
  static serializeBinaryToWriter(message: CreateSongResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSongResponse;
  static deserializeBinaryFromReader(message: CreateSongResponse, reader: jspb.BinaryReader): CreateSongResponse;
}

export namespace CreateSongResponse {
  export type AsObject = {
    ok: boolean,
  }
}

export class GetSongRequest extends jspb.Message {
  getSongId(): string;
  setSongId(value: string): GetSongRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSongRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetSongRequest): GetSongRequest.AsObject;
  static serializeBinaryToWriter(message: GetSongRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSongRequest;
  static deserializeBinaryFromReader(message: GetSongRequest, reader: jspb.BinaryReader): GetSongRequest;
}

export namespace GetSongRequest {
  export type AsObject = {
    songId: string,
  }
}

export class GetSongResponse extends jspb.Message {
  getSong(): Song | undefined;
  setSong(value?: Song): GetSongResponse;
  hasSong(): boolean;
  clearSong(): GetSongResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSongResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetSongResponse): GetSongResponse.AsObject;
  static serializeBinaryToWriter(message: GetSongResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSongResponse;
  static deserializeBinaryFromReader(message: GetSongResponse, reader: jspb.BinaryReader): GetSongResponse;
}

export namespace GetSongResponse {
  export type AsObject = {
    song?: Song.AsObject,
  }
}

export class UpdateSongRequest extends jspb.Message {
  getSongId(): string;
  setSongId(value: string): UpdateSongRequest;

  getSong(): Song | undefined;
  setSong(value?: Song): UpdateSongRequest;
  hasSong(): boolean;
  clearSong(): UpdateSongRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateSongRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateSongRequest): UpdateSongRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateSongRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateSongRequest;
  static deserializeBinaryFromReader(message: UpdateSongRequest, reader: jspb.BinaryReader): UpdateSongRequest;
}

export namespace UpdateSongRequest {
  export type AsObject = {
    songId: string,
    song?: Song.AsObject,
  }
}

export class UpdateSongResponse extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): UpdateSongResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateSongResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateSongResponse): UpdateSongResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateSongResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateSongResponse;
  static deserializeBinaryFromReader(message: UpdateSongResponse, reader: jspb.BinaryReader): UpdateSongResponse;
}

export namespace UpdateSongResponse {
  export type AsObject = {
    ok: boolean,
  }
}

export class DeleteSongRequest extends jspb.Message {
  getSongId(): string;
  setSongId(value: string): DeleteSongRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteSongRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteSongRequest): DeleteSongRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteSongRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteSongRequest;
  static deserializeBinaryFromReader(message: DeleteSongRequest, reader: jspb.BinaryReader): DeleteSongRequest;
}

export namespace DeleteSongRequest {
  export type AsObject = {
    songId: string,
  }
}

export class DeleteSongResponse extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): DeleteSongResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteSongResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteSongResponse): DeleteSongResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteSongResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteSongResponse;
  static deserializeBinaryFromReader(message: DeleteSongResponse, reader: jspb.BinaryReader): DeleteSongResponse;
}

export namespace DeleteSongResponse {
  export type AsObject = {
    ok: boolean,
  }
}

export class GetSongsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSongsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetSongsRequest): GetSongsRequest.AsObject;
  static serializeBinaryToWriter(message: GetSongsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSongsRequest;
  static deserializeBinaryFromReader(message: GetSongsRequest, reader: jspb.BinaryReader): GetSongsRequest;
}

export namespace GetSongsRequest {
  export type AsObject = {
  }
}

export class GetSongsResponse extends jspb.Message {
  getSongList(): Array<Song>;
  setSongList(value: Array<Song>): GetSongsResponse;
  clearSongList(): GetSongsResponse;
  addSong(value?: Song, index?: number): Song;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSongsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetSongsResponse): GetSongsResponse.AsObject;
  static serializeBinaryToWriter(message: GetSongsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSongsResponse;
  static deserializeBinaryFromReader(message: GetSongsResponse, reader: jspb.BinaryReader): GetSongsResponse;
}

export namespace GetSongsResponse {
  export type AsObject = {
    songList: Array<Song.AsObject>,
  }
}

export class PauseRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PauseRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PauseRequest): PauseRequest.AsObject;
  static serializeBinaryToWriter(message: PauseRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PauseRequest;
  static deserializeBinaryFromReader(message: PauseRequest, reader: jspb.BinaryReader): PauseRequest;
}

export namespace PauseRequest {
  export type AsObject = {
  }
}

export class PauseResponse extends jspb.Message {
  getOk(): boolean;
  setOk(value: boolean): PauseResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PauseResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PauseResponse): PauseResponse.AsObject;
  static serializeBinaryToWriter(message: PauseResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PauseResponse;
  static deserializeBinaryFromReader(message: PauseResponse, reader: jspb.BinaryReader): PauseResponse;
}

export namespace PauseResponse {
  export type AsObject = {
    ok: boolean,
  }
}

export class PlayRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlayRequest): PlayRequest.AsObject;
  static serializeBinaryToWriter(message: PlayRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayRequest;
  static deserializeBinaryFromReader(message: PlayRequest, reader: jspb.BinaryReader): PlayRequest;
}

export namespace PlayRequest {
  export type AsObject = {
  }
}

export class PlayResponse extends jspb.Message {
  getMusic(): Music | undefined;
  setMusic(value?: Music): PlayResponse;
  hasMusic(): boolean;
  clearMusic(): PlayResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PlayResponse): PlayResponse.AsObject;
  static serializeBinaryToWriter(message: PlayResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayResponse;
  static deserializeBinaryFromReader(message: PlayResponse, reader: jspb.BinaryReader): PlayResponse;
}

export namespace PlayResponse {
  export type AsObject = {
    music?: Music.AsObject,
  }
}

