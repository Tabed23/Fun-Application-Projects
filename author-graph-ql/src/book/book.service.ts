import { Injectable } from '@nestjs/common';
import { BookEntity } from './entity/book.entity';


@Injectable()
export class BookService {
    public booksData: BookEntity[] = [];

    addBook(book: BookEntity): String {
        this.booksData.push(book);
        return "book added successfully"
    }

    updateBook(id :number, updateBook: BookEntity): String{
        for (let x = 0; x < this.booksData.length; x++){
            if (this.booksData[x].id == id){
                this.booksData[x] = updateBook;
                return "book updated successfully"
            }
        }
        return "book not found"

    }

    deleteBook(id :number): String{
        this.booksData = this.booksData.filter((book) => book.id == id)
        return "book deleted successfully"
    }

    findBookById(id :number): BookEntity{
        for (let x = 0; x < this.booksData.length; x++){
            if (this.booksData[x].id == id){
                return this.booksData[x]
            }
        }
        return null
    }

    findAllBooks(): BookEntity[] {
        return this.booksData;
    }
}
