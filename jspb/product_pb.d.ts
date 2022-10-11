import * as jspb from 'google-protobuf'



export class ProductList extends jspb.Message {
  getListList(): Array<Product>;
  setListList(value: Array<Product>): ProductList;
  clearListList(): ProductList;
  addList(value?: Product, index?: number): Product;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProductList.AsObject;
  static toObject(includeInstance: boolean, msg: ProductList): ProductList.AsObject;
  static serializeBinaryToWriter(message: ProductList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProductList;
  static deserializeBinaryFromReader(message: ProductList, reader: jspb.BinaryReader): ProductList;
}

export namespace ProductList {
  export type AsObject = {
    listList: Array<Product.AsObject>,
  }
}

export class Product extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): Product;

  getName(): string;
  setName(value: string): Product;

  getLabelMap(): jspb.Map<string, string>;
  clearLabelMap(): Product;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Product.AsObject;
  static toObject(includeInstance: boolean, msg: Product): Product.AsObject;
  static serializeBinaryToWriter(message: Product, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Product;
  static deserializeBinaryFromReader(message: Product, reader: jspb.BinaryReader): Product;
}

export namespace Product {
  export type AsObject = {
    uuid: string,
    name: string,
    labelMap: Array<[string, string]>,
  }
}

