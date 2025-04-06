# SchoolSearch v2

# Important

### Use the following link to make your API key https://collegescorecard.ed.gov/data/api-documentation

## Whats New?
SchoolSearch2 is a new and improved take on my original School-Search. This new version aims to further expand acsess to college comparison tools to help students evaluate there options.
- made with golang: faster, like a lot faster. After a test of comparing SchoolSearch2 to its predecessor in a 3 period test I found that SchoolSearch2 was roughly 1.9x faster.
- improved interface: an easier to use interface.
- portable: no more need to play around with setting up python simply add your API key at the beginning  of the program and your all good, use the executables or if you want to hard code your key download the code and add it to the main file in less than a minute you wont have to worry about it again.
- better data: While the previous one relied on another website that had a risk of banning you for using it to web scrape while this one uses the US department of education score card for its data.



# School-Search v1
### A cli tool to quickly and easily generate a college spread sheet.

## Setup

- Requirements:
  - Python3: [Pythons Download Page](https://www.python.org/downloads/) or you could use it on the cloud with [replit.com](https://replit.com)
  - beautiful soup 4: `pip install beautifulsoup4`
  - lxml: `pip install lxml`
- How to Use:
    - `python3 schoolSearch.py` or `python schoolSearch.py`
    - then follow the prompts
    - make sure to use links from [https://www.appily.com/](https://www.appily.com/colleges)
