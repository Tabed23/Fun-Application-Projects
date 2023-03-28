import { CreateBookArgs, UpdateBookArgs } from './args/book.args';
import { Injectable } from '@nestjs/common';
import { BookEntity } from './entity/book.entity';
import { Repository } from 'typeorm';
import { InjectRepository } from '@nestjs/typeorm';


@Injectable()
export class BookService {

    constructor(@InjectRepository(BookEntity) public readonly bookRepo: Repository<BookEntity>){}

   async addBook(bookArgs: CreateBookArgs): Promise<String> {
    let book : BookEntity = new BookEntity();
    book.title = bookArgs.title;
    book.price = bookArgs.price;
    await this.bookRepo.save(book);
    return "book created successfully"
    }

    async updateBook(updateBook: UpdateBookArgs): Promise<String>{
        let book : BookEntity = await this.bookRepo.findOne({where: {id: updateBook.id}});
        book.title = updateBook.title;
        book.price = updateBook.price;
        await this.bookRepo.save(book);
        return "book updated successfully";
    }

    async deleteBook(id :number): Promise<String>{
        await this.bookRepo.delete(id)
        return "Book deleted"
    }

   async findBookById(id :number): Promise<BookEntity>{
    const book = await this.bookRepo.findOne({where: {id: id}});
    return book;
    }

    async findAllBooks(): Promise<BookEntity[]> {
        const books = await this.bookRepo.find();
        return books;
    }
}
