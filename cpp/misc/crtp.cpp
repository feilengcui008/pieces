#include <iostream>
#include <memory>

// crtp way
template <typename Derived>
class Parent {
   public:
    void talkInterface() {
        std::cout << "in crtp parent talk" << std::endl;
        static_cast<Derived *>(this)->talkImpl();
    }

   private:
    void talkImpl() { std::cout << "in crtp parent talk impl" << std::endl; }
};

class ChildFirst : public Parent<ChildFirst> {
   public:
    void talkImpl() {
        std::cout << "in crtp ChildFirst talk impl" << std::endl;
    }
};

class ChildSecond : public Parent<ChildSecond> {
   public:
    void talkImpl() {
        std::cout << "in crtp ChildSecond talk impl" << std::endl;
    }
};

int main(int argc, char *argv[]) {
    std::shared_ptr<ChildFirst> ptr1 =
        std::make_shared<ChildFirst>(ChildFirst());
    std::shared_ptr<ChildSecond> ptr2 =
        std::make_shared<ChildSecond>(ChildSecond());
    ptr1->talkInterface();
    ptr2->talkInterface();
    return 0;
}
