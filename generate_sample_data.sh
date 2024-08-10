#!/bin/bash

# Create directory structure
mkdir -p personal_data/{journals,blog_posts,social_media,references}

# Function to generate random tags
generate_tags() {
    tags=("#AI" "#personalGrowth" "#technology" "#learning" "#productivity" "#health" "#creativity" "#mindfulness")
    num_tags=$((RANDOM % 3 + 1))
    selected_tags=""
    for i in $(seq 1 $num_tags); do
        random_tag=${tags[$RANDOM % ${#tags[@]}]}
        selected_tags="$selected_tags $random_tag"
    done
    echo $selected_tags
}

# Function to generate dates
generate_date() {
    days_ago=$1
    date -d "$days_ago days ago" +%Y-%m-%d
}

# Generate journal entries
for i in {1..10}; do
    date=$(generate_date $i)
    tags=$(generate_tags)
    echo "Dear Diary,

Today I reflected on my personal growth journey. $tags

I've been thinking about how I can improve my skills and become more productive. 
There's so much to learn, especially in the field of AI and technology.

I'm excited about the possibilities and looking forward to what tomorrow brings." > personal_data/journals/${date}_daily_thoughts.txt
done

# Generate blog posts
for i in {1..5}; do
    date=$(generate_date $((i*30)))
    tags=$(generate_tags)
    echo "The Future of AI: Possibilities and Challenges

$tags

Artificial Intelligence is rapidly evolving, opening up new possibilities across various industries. 
From healthcare to finance, AI is transforming the way we work and live.

However, with great power comes great responsibility. We must consider the ethical implications 
of AI and ensure that it's developed and used in ways that benefit humanity as a whole.

What are your thoughts on the future of AI?" > personal_data/blog_posts/${date}_ai_thoughts.txt
done

# Generate social media posts
for i in {1..15}; do
    date=$(generate_date $i)
    tags=$(generate_tags)
    echo "Just finished reading an fascinating article on machine learning applications in healthcare. 
The potential for AI to improve patient outcomes is truly exciting! $tags 
#thoughtLeadership" > personal_data/social_media/${date}_twitter_post.txt
done

# Generate references
for i in {1..7}; do
    date=$(generate_date $((i*30)))
    tags=$(generate_tags)
    echo "Title: 'The Singularity Is Near' by Ray Kurzweil

$tags

Key takeaways:
1. The pace of technological progress is exponential, not linear.
2. AI will surpass human intelligence in the near future.
3. The fusion of human and machine intelligence will lead to unprecedented advancements.

Thoughts: Kurzweil's predictions are bold and thought-provoking. While some may seem far-fetched, 
his track record of accurate predictions makes this book a must-read for anyone interested in 
the future of technology and humanity." > personal_data/references/${date}_book_notes.txt
done

echo "Sample data generation complete."