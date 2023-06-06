<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<div align="left">

<h1>Dolorem</h1>
  <p>
    A highly inclusive & customizable Lorem Ipsum generator for Go.
    <br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About this module

Generate random Lorem Ipsum text for your project, drawing words from the full Latin dictionary!

This module provides Lorem in the following formats:

  * Paragraph(s) - Random paragraph, with customizable options for number of paragraphs, number of sentences per paragraph & number of words per sentence.
  * Sentence - Random sentence, with customizable options for number of words per sentence.
  * Word - A random word.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE -->
## Usage

#### Import the module into your project:

~~~
package myAmazingPackage

import (
  // ...
  "github.com/jfajardo5/dolorem"
)
~~~

#### Initialize Dolorem, and code away!

~~~
//***************
// Default Options
//*****************

// Initialize Dolorem
lorem := dolorem.Ipsum()

// Random Paragraph
fmt.Println(lorem.Paragraph())

// Random Sentence
fmt.Println(lorem.Sentence())

// Random Word
fmt.Println(lorem.Word())
~~~

#### This module provides customizable options for specifying:

  * Number of paragraphs
  * Number of sentences per paragraph
  * Number of words per sentence

~~~
//**************
// Custom Options
//****************

// Single Sentence with x number of Words
fmt.Println(lorem.Sentence(5))

// x number of Paragraphs
fmt.Println(lorem.Paragraph(3))

// x number of Paragraphs with y number of Sentences per Paragraph
fmt.Println(lorem.Paragraph(2, 10))

// x number of Paragraphs with y number of Sentences with z number of Words per Sentence
fmt.Println(lorem.Paragraph(2, 12, 6))
~~~

#### By default, the very first Paragraph in a set will always start with "Dolorem ipsum dolor sit amet,"
#### This default can be overridden to your liking:

~~~
// Initialize Dolorem
lorem := dolorem.Ipsum()

// Override ParagraphStarter
lorem.ParagraphStarter = "My custom paragraph starter,"

// Random Paragraph with custom ParagraphStarter
fmt.Println(lorem.Paragraph())

// It can also be overridden with an empty string
lorem.ParagraphStarter = ""
  
// Random Paragraph with no ParagraphStarter
fmt.Println(lorem.Paragraph())
~~~

#### You may also override the included dictionary with your custom dictionary.

#### It can adapt any other []string!

~~~
// Initialize Dolorem
lorem := dolorem.Ipsum()

// Override dictionary
lorem.Dictionary = []string{"my", "custom", "dictionary"}

// Random Paragraph from custom dictionary
 fmt.Println(lorem.Paragraph())
~~~

#### For even more advanced usage, you are able to override the default seed:

~~~
// Initialize Dolorem
lorem := dolorem.Ipsum()

// Override Seed
lorem.Seed = rand.New(rand.NewSource(time.Now().Unix()))
~~~

#### Enjoy!



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Julio Fajardo - [jfajardo.net](https://jfajardo.net)

Project Link: [https://github.com/jfajardo5/dolorem](https://github.com/jfajardo5/dolorem)

<p align="right">(<a href="#readme-top">back to top</a>)</p>




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/jfajardo5/dolorem.svg?style=for-the-badge
[contributors-url]: https://github.com/jfajardo5/dolorem/graphs/contributors
[issues-shield]: https://img.shields.io/github/issues/jfajardo5/dolorem.svg?style=for-the-badge
[issues-url]: https://github.com/jfajardo5/dolorem/issues
[license-shield]: https://img.shields.io/github/license/jfajardo5/dolorem.svg?style=for-the-badge
[license-url]: https://github.com/jfajardo5/dolorem/blob/main/LICENSE
