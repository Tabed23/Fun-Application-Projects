
/*
 * -------------------------------------------------------
 * THIS FILE WAS AUTOMATICALLY GENERATED (DO NOT MODIFY)
 * -------------------------------------------------------
 */

/* tslint:disable */
/* eslint-disable */

export interface CreateBookArgs {
    title: string;
    price: number;
}

export interface UpdateBookArgs {
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
    createBook(addBookArgs: CreateBookArgs): string | Promise<string>;
    UpdateBook(updateBookArgs: UpdateBookArgs): string | Promise<string>;
}

type Nullable<T> = T | null;
