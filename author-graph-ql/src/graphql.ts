
/*
 * -------------------------------------------------------
 * THIS FILE WAS AUTOMATICALLY GENERATED (DO NOT MODIFY)
 * -------------------------------------------------------
 */

/* tslint:disable */
/* eslint-disable */

export interface BookArgs {
    id: number;
    title: string;
    price: number;
}

export interface Book {
    id: number;
    title: string;
    price: number;
}

export interface IQuery {
    index(): string | Promise<string>;
    getAllBooks(): Book[] | Promise<Book[]>;
    getBookByID(bookID: number): Nullable<Book> | Promise<Nullable<Book>>;
}

export interface IMutation {
    deleteBookByID(bookID: number): string | Promise<string>;
    createBook(addBookArgs: BookArgs): string | Promise<string>;
    UpdateBook(bookID: number, updateBookArgs: BookArgs): string | Promise<string>;
}

type Nullable<T> = T | null;
