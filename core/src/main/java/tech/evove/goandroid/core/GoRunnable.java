package tech.evove.goandroid.core;

final class GoRunnable implements core.Runnable, Runnable {
    private final Runnable actual;

    private GoRunnable(Runnable actual) {
        this.actual = actual;
    }

    static core.Runnable wrap(Runnable runnable) {
        if (runnable instanceof core.Runnable) {
            return (core.Runnable) runnable;
        }
        return new GoRunnable(runnable);
    }

    @Override
    public void run() {
        actual.run();
    }
}
